// Code generated by "golang-cloudformation". DO NOT EDIT.

package {{ .GoPackageName }}

{{ if .GoImports }}
    import (
        {{- range .GoImports -}}
            "{{ . }}"
        {{ end }}
    )
{{ end }}

{{ if .IsTopLevelResourceType }}
var (
    _ cfz.ResourcePartialLogicalName = (*{{ .GoName }})(nil)
    _ cfz.Resource = (*{{ .GoName }})(nil)
)
{{ end }}

const (
    {{ .GoName }}__Type = "{{ .Name }}"
)

{{ if .Attributes }}
var (
    {{ .GoName }}__AttributesMap = struct {
        {{ range $k, $v := .Attributes -}}
            {{ $v.GoName }} string
        {{ end }}
    }{
        {{ range $k, $v := .Attributes -}}
            {{ $v.GoName }}: "{{ $v.Name }}",
        {{ end }}
    }

    {{ .GoName }}__AttributesSlice = []string{
        {{ range $k, $v := .Attributes -}}
            {{ $.GoName }}__AttributesMap.{{ $v.GoName }},
        {{ end }}
    }
)
{{ end }}

{{ if .Properties }}
var (
    {{ .GoName }}__PropertiesMap = struct {
        {{ range $k, $v := .Properties -}}
            {{ $v.GoName }} string
        {{ end }}
    }{
        {{ range $k, $v := .Properties -}}
            {{ $v.GoName }}: "{{ $v.Name }}",
        {{ end }}
    }

    {{ .GoName }}__PropertiesSlice = []string{
        {{ range $k, $v := .Properties -}}
            {{ $.GoName }}__PropertiesMap.{{ $v.GoName }},
        {{ end }}
    }
)
{{ end }}

// {{ .GoName }} is a binding for {{ .Name }}.
// See: {{ .DocumentationURL }}
type {{ .GoName }} struct {
    {{ if .IsTopLevelResourceType -}}
        // CF_LogicalName is the CloudFormation logical name for this resource in the template.
        CF_LogicalName string `json:"-"`

        // CF_DependsOn indicates which resources must be created before this one.
        CF_DependsOn []cfz.ResourcePartialLogicalName `json:"-"`
    {{ end }}

    {{- range $k, $v := .Properties }}
        // {{ $v.Name }} is a resource property.
        // See: {{ $v.DocumentationURL }}
        {{ $v.GoName }} {{ $v.GoType }} `json:"{{ $v.Name }},omitempty"`
    {{ end }}
}

{{ if .IsTopLevelResourceType }}
    // GetLogicalName returns the CloudFormation logical name for this resource in the template.
    // It implements the cfz.Resource and cfz.ResourcePartialLogicalName interface.
    func (v *{{ .GoName }}) GetResourceLogicalName() string {
        return v.CF_LogicalName
    }
{{ end }}

// GetType returns the CloudFormation type.{{ if .IsTopLevelResourceType }}
// It implements the cfz.Resource interface.{{ end }}
func (v *{{ .GoName }}) GetType() string {
    return {{ .GoName }}__Type
}

{{ if .IsTopLevelResourceType }}
    // Ref returns an Expression[string] that resolves to the Ref intrinsic function for this resource.
    func (v *{{ .GoName }}) Ref() cfz.Expression[string] {
        return cfz.Ref(cfz.V(v.GetResourceLogicalName()))
    }

    {{ range $k, $v := .Attributes }}
        // {{ $v.SupportGetAttFunctionName }}__{{ $v.GoName }} returns an XXX[{{ $v.GoGenericType }}] that resolves to the FN::GetAtt intrinsic function for this resource.
        // Attribute: {{ $v.Name }}
        func (v *{{ $.GoName }}) {{ $v.SupportGetAttFunctionName }}__{{ $v.GoName }}() {{ $v.GoType }} {
            return cfz.{{ $v.SupportGetAttFunctionName }}[{{ $v.GoGenericType }}](cfz.V(v.GetResourceLogicalName()), cfz.V({{ $.GoName }}__AttributesMap.{{ $v.GoName }}))
        }
    {{ end }}

    // MarshalJSON implements the cfz.Resource interface.
    func (v *{{ .GoName }}) MarshalJSON() ([]byte, error) {
        if v == nil {
            return []byte(`null`), nil
        }

        type CF_Properties {{ .GoName }}

        return json.Marshal(struct {
            Type string `json:"Type"`
            DependsOn []string `json:"DependsOn,omitempty"`
            Properties *CF_Properties `json:"Properties,omitempty"`
        }{
            Type: v.GetType(),
            DependsOn: v.getDependsOn(),
            Properties: (*CF_Properties)(v),
        })
    }

    func (v *{{ .GoName }}) getDependsOn() []string {
        dependsOn := make([]string, 0, len(v.CF_DependsOn))

        for _, r := range v.CF_DependsOn {
            dependsOn = append(dependsOn, r.GetResourceLogicalName())
        }

        return dependsOn
    }
{{ end }}