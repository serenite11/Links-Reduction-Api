<h1>Links-Reduction-Api</h1>

<h3>Для запуска сервиса клонируйте репозиторий и выполните команду:</h3>
<h3><code>make compose</code></h3>
<h2>REST:</h2>

<h2><code>POST</code></h2>

<h3><code>http://localhost:8080/ </code></h3>  

<h3><code>&nbsp;&nbsp;{<br>&nbsp; &nbsp;&nbsp;&nbsp;"long_url":"https://www.youtube.com" <br>&nbsp; &nbsp;}</code></h3>

<h2><code>Response</code></h2>
<h3><code>&nbsp;&nbsp;{<br>&nbsp; &nbsp;&nbsp;&nbsp;"shortLink":"aB3jOoPlM_" <br>&nbsp;&nbsp; }</code></h3>

<h2><code>GET</code></h2>

<h3><code>http://localhost:8080/ </code></h3>

<h3><code>&nbsp;&nbsp;{<br>&nbsp; &nbsp;&nbsp;&nbsp;"short_url":"aB3jOoPlM_"<br>&nbsp;&nbsp;&nbsp;}</code></h3>

<h2><code>Response</code></h2>
<h3><code>&nbsp;&nbsp;{<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"longLink":"https://www.youtube.com" <br>&nbsp;&nbsp; }</code></h3>

<h2>Grpc</h2>

<h3><code>localhost:5500</code></h3>

<h2><code>CreateShortUrl</code></h2>

<h3><code>&nbsp;&nbsp;{<br>&nbsp; &nbsp;&nbsp;&nbsp;"url":"https://www.youtube.com"<br> &nbsp;&nbsp; } </code></h3>
<h2><code>Response</code></h2>
<h3><code>&nbsp;&nbsp;{<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"url":"aB3jOoPlM_" <br>&nbsp;&nbsp; }</code></h3>
<h2><code>GetLongUrl</code></h2>

<h3><code>&nbsp;&nbsp;{<br>&nbsp; &nbsp;&nbsp;&nbsp;"url":"aB3jOoPlM_"<br> &nbsp;&nbsp; } </code></h3>
<h2><code>Response</code></h2>
<h3><code>&nbsp;&nbsp;{<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"url":"https://www.youtube.com" <br>&nbsp;&nbsp; }</code></h3>

<h3>Для тестов используйте команду:</h3>
<h3><code>make test</code></h3>


