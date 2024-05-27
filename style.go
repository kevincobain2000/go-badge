package badge

var flatTemplate = stripXmlWhitespace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20">
  <linearGradient id="smooth" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".0"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>

  <mask id="round">
    <rect width="{{.Bounds.Dx}}" height="20" rx="3" fill="{{or .BadgeColor "#fff" | html}}"/>
  </mask>

  <g mask="url(#round)">
    <rect width="{{.Bounds.SubjectDx}}" height="20" fill="#555"/>
    <rect x="{{.Bounds.SubjectDx}}" width="{{.Bounds.StatusDx}}" height="20" fill="{{or .Color "#4c1" | html}}"/>
    <rect width="{{.Bounds.Dx}}" height="20" fill="url(#smooth)"/>
  </g>

  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="{{.Bounds.SubjectX}}" y="15" fill="#010101" fill-opacity=".3">{{.Subject | html}}</text>
    <text x="{{.Bounds.SubjectX}}" y="14">{{.Subject | html}}</text>
    <text x="{{.Bounds.StatusX}}" y="15" fill="#010101" fill-opacity=".3">{{.Status | html}}</text>
    <text x="{{.Bounds.StatusX}}" y="15" fill="{{or .LabelColor "#fff" | html}}">{{.Status | html}}</text>
  </g>
</svg>
`)

var flatSocialTemplate = stripXmlWhitespace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Bounds.Dx}}" height="20">
	<style>a:hover #llink{fill:url(#b);stroke:#ccc}a:hover #rlink{fill:#4183c4}</style>
	<linearGradient id="a" x2="0" y2="100%">
		<stop offset="0" stop-color="#fcfcfc" stop-opacity="0"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<linearGradient id="b" x2="0" y2="100%">
		<stop offset="0" stop-color="#ccc" stop-opacity=".1"/>
		<stop offset="1" stop-opacity=".1"/>
	</linearGradient>
	<g stroke="#d5d5d5">
		<rect stroke="none" fill="#fcfcfc" x="0.5" y="0.5" width="{{.Bounds.Dx}}" height="19" rx="2"/>
		<rect x="{{.Bounds.SubjectDx}}" y="0.5" width="{{.Bounds.StatusDx}}" height="19" rx="2" fill="{{or .Color "#fafafa" | html}}"/>
		<rect x="{{.Bounds.SubjectDx}}" y="7.5" width="0.5" height="5" stroke="#fafafa"/>
		<path d="M{{.Bounds.SubjectDx}} 6.5 l-3 3v1 l3 3" stroke="#d5d5d5" fill="#fafafa"/>
	</g>
	<g aria-hidden="true" fill="#333" text-anchor="middle" font-family="Helvetica Neue,Helvetica,Arial,sans-serif" text-rendering="geometricPrecision" font-weight="700" font-size="11px" line-height="14px">
		<rect id="llink" stroke="#d5d5d5" fill="url(#a)" x=".5" y=".5" width="{{.Bounds.Dx}}" height="19" rx="2"/>
		<text x="{{.Bounds.SubjectX}}" y="15" fill="#010101" fill-opacity=".3">{{.Subject | html}}</text>
		<text x="{{.Bounds.SubjectX}}" y="14">{{.Subject | html}}</text>
		<text x="{{.Bounds.StatusX}}" y="15" fill="#010101" fill-opacity=".3">{{.Status | html}}</text>
		<text id="rlink" x="{{.Bounds.StatusX}}" y="15" fill="{{or .LabelColor "#fff" | html}}">{{.Status | html}}</text>
	</g>
</svg>
`)
