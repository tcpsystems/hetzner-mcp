package server

import (
	"context"
	"strconv"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MonthlyBudgetGross is the standing budget ceiling for this project, in the
// account's currency. Below this is reasonable; above it needs a decision
// from the human operator, not a seat. See plan/projects/hetzner-finance.md.
const MonthlyBudgetGross = 100.0

type resourceCost struct {
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Type         string `json:"type,omitempty"`
	Location     string `json:"location,omitempty"`
	Status       string `json:"status,omitempty"`
	MonthlyGross string `json:"monthly_gross,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Idle         bool   `json:"idle"`
	IdleReason   string `json:"idle_reason,omitempty"`
}

type costsOut struct {
	Currency              string         `json:"currency"`
	BudgetMonthlyGross    float64        `json:"budget_monthly_gross"`
	EstimatedMonthlyGross float64        `json:"estimated_monthly_gross"`
	OverBudget            bool           `json:"over_budget"`
	Resources             []resourceCost `json:"resources"`
	IdleResources         []resourceCost `json:"idle_resources,omitempty"`
	Hint                  string         `json:"hint"`
}

// getCosts reports live spend against the standing budget by combining the
// account's resource inventory with Hetzner's own /pricing rates. It never
// changes a resource; it only reads.
func (a *App) getCosts(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, costsOut, error) {
	client, err := a.api()
	if err != nil {
		return nil, costsOut{}, err
	}

	pricing, _, err := client.Pricing.Get(ctx)
	if err != nil {
		return nil, costsOut{}, err
	}
	rateByType := map[string]string{}
	for _, st := range pricing.ServerTypes {
		if len(st.Pricings) == 0 {
			continue
		}
		rateByType[st.ServerType.Name] = st.Pricings[0].Monthly.Gross
	}

	servers, _, err := client.Server.List(ctx, hcloud.ServerListOpts{})
	if err != nil {
		return nil, costsOut{}, err
	}

	out := costsOut{
		Currency:           pricing.Currency,
		BudgetMonthlyGross: MonthlyBudgetGross,
		Hint: "Resources here cost money every month whether or not they are used. " +
			"idle=true means this resource is not doing known work (no server to anchor it) and is a " +
			"candidate to delete, not just note. Do not create a new server to replace one already idle.",
	}

	var total float64
	for _, s := range servers {
		typeName := ""
		if s.ServerType != nil {
			typeName = s.ServerType.Name
		}
		gross := rateByType[typeName]
		if f, err := strconv.ParseFloat(gross, 64); err == nil {
			total += f
		}
		loc := ""
		if s.Location != nil {
			loc = s.Location.Name
		}
		out.Resources = append(out.Resources, resourceCost{
			Kind:         "server",
			Name:         s.Name,
			Type:         typeName,
			Location:     loc,
			Status:       string(s.Status),
			MonthlyGross: gross,
			Currency:     pricing.Currency,
		})
	}

	// Volumes, floating IPs, and load balancers bill even when detached or
	// unattached. List them and flag any with no server, since those are
	// the unutilized-infrastructure case the operator asked to avoid.
	volumes, _, err := client.Volume.List(ctx, hcloud.VolumeListOpts{})
	if err == nil {
		for _, v := range volumes {
			idle := v.Server == nil
			rc := resourceCost{
				Kind:     "volume",
				Name:     v.Name,
				Location: v.Location.Name,
				Status:   string(v.Status),
				Idle:     idle,
			}
			if idle {
				rc.IdleReason = "not attached to any server"
				out.IdleResources = append(out.IdleResources, rc)
			}
			out.Resources = append(out.Resources, rc)
		}
	}

	lbs, _, err := client.LoadBalancer.List(ctx, hcloud.LoadBalancerListOpts{})
	if err == nil {
		for _, lb := range lbs {
			idle := len(lb.Targets) == 0
			rc := resourceCost{
				Kind:   "load_balancer",
				Name:   lb.Name,
				Status: "unknown",
				Idle:   idle,
			}
			if idle {
				rc.IdleReason = "no targets configured"
				out.IdleResources = append(out.IdleResources, rc)
			}
			out.Resources = append(out.Resources, rc)
		}
	}

	out.EstimatedMonthlyGross = total
	out.OverBudget = total > MonthlyBudgetGross
	if out.OverBudget {
		out.Hint = "ESTIMATED SPEND IS OVER THE £100 STANDING BUDGET. " + out.Hint
	}
	return nil, out, nil
}
