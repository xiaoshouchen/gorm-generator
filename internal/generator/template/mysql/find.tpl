{{define "find"}}
    {{if withCache}}
        {{template "findByPkWithCache" .}}
        {{template "findByPksWithCache" .}}
        {{template "findByUniqueWithCache" .}}
        {{template "findByUniques" .}}
        {{template "findByIndex" .}}
    {{else}}
        {{template "findByPk" .}}
        {{template "findByPks" .}}
        {{template "findByUnique" .}}
        {{template "findByUniques" .}}
        {{template "findByIndex" .}}
    {{end}}
{{end}}