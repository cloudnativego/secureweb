package server

import (
	"net/http"
	"text/template"
)

var bodyTemplate = `
<script src="https://cdn.auth0.com/js/lock-8.2.min.js"></script>
<script type="text/javascript">

  var lock = new Auth0Lock('{{.ClientID}}', '{{.Domain}}');


  function signin() {
    lock.show({
        callbackURL: '{{.CallbackURL}}'
      , responseType: 'code'
      , authParams: {
        scope: 'openid profile'
      }
    });
  }
</script>
<button onclick="window.signin();">Login</button>
`

func homeHandler(config *authConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("htmlz").Parse(bodyTemplate))
		t.Execute(w, config)
	}
}
