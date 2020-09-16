package common

import (
	"fmt"
	"strings"
)

func NewFilter() *Filter {
	return &Filter{
		NextQueryParams: new(ParamCollection),
	}
}
func (f Filter) BuildFilter() (r string) {
	f.NextQueryParams.Params = make(map[string][]string)
	var filter string
	var qAnd string
	var subStr string
	if len(f.URLQueryParams) > 0 {
		i := 0
		for k, v := range f.URLQueryParams {
			if k == "limit" || k == "offset" || k == "orderCol" || k == "orderDirection" {
				continue
			}
			f.NextQueryParams.Params[k] = v
			subStr = ""
			qOr := ""
			for ii := 0; ii < len(v); ii++ {
				if ii != 0 {
					qOr = " OR "
				}
				searchTerm := f.FormatURLQry(k, v[ii])
				if searchTerm != "" {
					subStr = subStr + qOr + searchTerm
				}
			}
			subStr = "(" + subStr + ")"
			if i != 0 {
				qAnd = " AND "
			}
			if len(subStr) > 2 {
				filter = filter + qAnd + subStr
			}
			i++
		}
		if filter != "" {
			r = " WHERE " + filter
		}
	}
	return
}

// FormatURLQry converts a field+value pair to a SQL WHERE clause.
func (f Filter) FormatURLQry(field string, val string) string {
	var resp string
	var eval string
	field = strings.ToLower(field)
	val = strings.TrimSpace(val)
	if strings.Contains(val, "*") {
		val = strings.Replace(val, "*", "%", -1)
		eval = " like '" + val + "'"
	} else {
		eval = " = '" + val + "'"
	}
	switch field {
	case "source":
		resp = field + eval
	case "status":
		resp = `status` + eval
	case "enabled":
		resp = `data->>'enabled'` + eval
	case "id":
		resp = field + eval
	case "last_modified":
		resp = field + eval
	case "md5hash":
		resp = field + eval
	case "platform":
		resp = `load_balancer->>'mfr'` + eval
	case "load_balancer_ip":
		resp = field + eval
	case "load_balancer":
		resp = field + eval
	case "dns":
		resp = `regexp_replace(regexp_replace(regexp_replace(data->>'dns','\[','{'),'\]','}'),'("|\s)','','g') like '%` + val + `%'`
	case "certificates":
		resp = `regexp_replace(regexp_replace(regexp_replace(data->>'certificates','\[','{'),'\]','}'),'("|\s)','','g') like '%` + val + `%'`
	case "pools":
		resp = `regexp_replace(regexp_replace(regexp_replace(data->>'pools','\[','{'),'\]','}'),'("|\s)','','g')  like '%` + val + `%'`
	case "ports":
		resp = `regexp_replace(regexp_replace(regexp_replace(data->>'ports','\[','{'),'\]','}'),'("|\s)','','g') like '%` + val + `%'`
	case "networksecuritypolicyname":
		resp = `data->'networksecuritypolicy'->>'name'` + eval
	case "networksecuritypolicynameenable":
		resp = `data @> '{"networksecuritypolicy":{"rules":[{"enable":` + val + `}]}}'`
	case "networksecuritypolicyaction":
		resp = `data @> '{"networksecuritypolicy":{"rules":[{"action":"` + val + `"}]}}'`
	case "networksecuritypolicyuuid":
		resp = `data->'networksecuritypolicy'->>'uuid'` + eval
	default:
		resp = `data->>'` + field + `'` + eval
	}
	return resp
}
func (f Filter) SetOrderBy(in string) (r string) {
	switch in {
	case "load_balancer_ip":
		r = "load_balancer_ip"
	case "_last_30":
		r = "data->>'_last_30'"
	case "cluster_ip":
		r = "data->>'cluster_ip'"
	case "cluster_dns":
		r = "data->>'cluster_dns'"
	case "mfr":
		r = "data->>'mfr'"
	case "model":
		r = "data->>'model'"
	case "ip":
		r = "data->>'ip'"
	case "name":
		r = "data->>'name'"
	case "product_code":
		r = "data->>'product_code'"
	case "service_type":
		r = "data->>'service_type'"
	case "status":
		r = "status"
	case "enabled":
		r = "data->>'enabled'"
	case "platform":
		r = "load_balancer->>'mfr'"
	default:
		r = "data->>'load_balancer_ip'"
	}
	return
}

// BuildSQLStmt generates a SQL statment using URL Params provided by the.
// Filter and DbTable objects.
func (f Filter) BuildSQLStmt() (r string, err error) {
	var limit string
	var offset string
	orderDirection := "asc"
	orderBy := "order by data->>'product_code' asc"
	whereClause := f.BuildFilter()
	if len(f.URLQueryParams["limit"]) == 1 {
		limit = fmt.Sprintf(" LIMIT %s ", f.URLQueryParams["limit"][0])
	}
	if len(f.URLQueryParams["offset"]) == 1 {
		offset = fmt.Sprintf(" OFFSET %s ", f.URLQueryParams["offset"][0])
	}
	if len(f.URLQueryParams["orderCol"]) == 1 {
		if len(f.URLQueryParams["orderDirection"]) == 1 {
			direction := strings.ToLower(f.URLQueryParams["orderDirection"][0])
			switch direction {
			case "asc":
				orderDirection = "asc"
			case "desc":
				orderDirection = "desc"
			default:
				orderDirection = "asc"
			}
		}
		orderBy = fmt.Sprintf(" ORDER BY %s %s", f.SetOrderBy(f.URLQueryParams["orderCol"][0]), orderDirection)
	}
	r = fmt.Sprintf(`SELECT * FROM
	(
		SELECT 
			a.id, 
			a.data, 
			a.last_modified, 
			a.md5hash, 
			a.load_balancer_ip, 
			a.last_modified_by, 
			a.load_balancer, 
			a.source, 
			CASE
				WHEN d.short is null THEN 'deployed'
				ELSE d.short 
			END as status 
		FROM public.%s as a 
		LEFT JOIN public.status as s
		ON a.id=s.id
		LEFT JOIN public.statusdescription as d  
		ON s.status_id=d.id
		) as d %s %s %s %s`, f.Table, whereClause, orderBy, limit, offset)
	// Convert sql to all lowercase.
	//r = strings.ToLower(r)
	return
}

// BuildVsSQLStmt generates a SQL statment using URL Params provided by the.
// Filter and DbTable objects.
func (f Filter) BuildVsSQLStmt() (r string, err error) {
	var limit string
	var offset string
	orderDirection := "asc"
	orderBy := "order by data->>'product_code' asc"
	whereClause := f.BuildFilter()
	if len(f.URLQueryParams["limit"]) == 1 {
		limit = fmt.Sprintf(" LIMIT %s ", f.URLQueryParams["limit"][0])
	}
	if len(f.URLQueryParams["offset"]) == 1 {
		offset = fmt.Sprintf(" OFFSET %s ", f.URLQueryParams["offset"][0])
	}
	if len(f.URLQueryParams["orderCol"]) == 1 {
		if len(f.URLQueryParams["orderDirection"]) == 1 {
			direction := strings.ToLower(f.URLQueryParams["orderDirection"][0])
			switch direction {
			case "asc":
				orderDirection = "asc"
			case "desc":
				orderDirection = "desc"
			default:
				orderDirection = "asc"
			}
		}
		orderBy = fmt.Sprintf(" ORDER BY %s %s", f.SetOrderBy(f.URLQueryParams["orderCol"][0]), orderDirection)
	}
	r = fmt.Sprintf(`SELECT 
	data->>'name' as name,
	data->>'ip' as ip,
	load_balancer_ip,
	data->>'platform' as platform,
	data->>'service_type' as service_type
	FROM
	(
		SELECT 
			a.id, 
			a.data, 
			a.last_modified, 
			a.md5hash, 
			a.load_balancer_ip, 
			a.last_modified_by, 
			a.load_balancer, 
			a.source, 
			CASE
				WHEN d.short is null THEN 'deployed'
				ELSE d.short 
			END as status 
		FROM public.%s as a 
		LEFT JOIN public.status as s
		ON a.id=s.id
		LEFT JOIN public.statusdescription as d  
		ON s.status_id=d.id
		WHERE a.data->>'ip' != '0.0.0.0'
		) as d %s %s %s %s`, f.Table, whereClause, orderBy, limit, offset)
	// Convert sql to all lowercase.
	//r = strings.ToLower(r)
	return
}

// BuildCountStmt generates a SQL statment using URL Params provided by the.
// Filter and DbTable objects.
func (f Filter) BuildCountStmt() (r string, err error) {
	whereClause := f.BuildFilter()
	r = fmt.Sprintf(`SELECT count(*) as total FROM
	(
		SELECT 
			a.id, 
			a.data, 
			a.last_modified, 
			a.md5hash, 
			a.load_balancer_ip, 
			a.last_modified_by, 
			a.load_balancer, 
			a.source, 
			CASE
				WHEN d.short is null THEN 'deployed'
				ELSE d.short 
			END as status 
		FROM public.%s as a 
		LEFT JOIN public.status as s
		ON a.id=s.id
		LEFT JOIN public.statusdescription as d  
		ON s.status_id=d.id
		) as d %s`, f.Table, whereClause)
	// Convert sql to all lowercase.
	//r = strings.ToLower(r)
	return
}
func (f Filter) SetQueryParams() (r string) {
	m := []string{}
	for k, v := range f.NextQueryParams.Params {
		for _, vv := range v {
			m = append(m, fmt.Sprintf("%s=%s", k, vv))
		}
	}
	r = strings.Join(m, "&")
	if r != "" {
		r = fmt.Sprintf("%s&", r)
	}
	return
}
