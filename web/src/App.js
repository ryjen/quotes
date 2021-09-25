import logo from './logo.png';
import '../node_modules/chota/dist/chota.css'
import './App.css';
import { Link } from 'react-router-dom'

function App() {
  return (

    <div className="App">
      <header id="header">
        <nav class="nav">
          <div class="nav-left">

            <a class="brand" href="/">
              <img src={logo} className="App-logo" alt="logo" />
              Parrot
            </a>

          </div>

          <div class="nav-right">
            <div class="tabs">
              <Link to="/app">App</Link>
              <a href="/docs">Docs</a>
              <a href="/about">About</a>
            </div>
          </div>
        </nav>
      </header>

      <div class="container" id="content">
        <div id="home">
          <h2>A colorful quotes API.</h2>
        </div>

        <div class="example">


        </div>


      </div>

      <footer id="footer">
        <small>
          all rights reserved &amp; &copy; 2021 by Micrantha Software
        </small>
      </footer>
    </div>
  );
}

export default App;
