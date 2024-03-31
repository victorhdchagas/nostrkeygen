import robotLogo from './assets/robot.svg'
import './App.css'
import KeyGeneratorForm from './components/KeyGeneratorForm'

function App() {
  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={robotLogo} className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>Nostr Key Generator</h1>
      <div className="card">
        <KeyGeneratorForm />
      </div>
    </>
  )
}

export default App
