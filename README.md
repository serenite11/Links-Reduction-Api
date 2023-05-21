<h2>Links-Reduction-Api</h2>

<h4>Для запуска сервиса клонируйте репозиторий и выполните команду:</h4>
<code>make compose</code>
<h2>REST:</h2>

<h2><code>POST</code></h2>

<code>http://localhost:8080/ </code> 

<code>{"long_url":"https://www.youtube.com"}</code>

<h4><code>Response</code></h4>
<code>{ "shortLink":"aB3jOoPlM_" }</code>

<h2><code>GET</code></h2>

<h3><code>http://localhost:8080/ </code></h3>

<code>{"short_url":"aB3jOoPlM_"}</code>

<h4><code>Response</code></h4>
<code>{"longLink":"https://www.youtube.com"}</code>

<h2>Grpc</h2>

<h3><code>localhost:5500</code></h3>

<code>CreateShortUrl</code><

<code>{"url":"https://www.youtube.com"} </code>
<h4><code>Response</code></h4>
<code>{"url":"aB3jOoPlM_"}</code>
<code>GetLongUrl</code>

<code>{"url":"aB3jOoPlM_"} </code>
<h4><code>Response</code></h4>
<code>{"url":"https://www.youtube.com" }</code></h3>

<h3>Для тестов используйте команду:</h3>
<code>make test</code>


