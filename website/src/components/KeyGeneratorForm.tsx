import { useState } from 'react'
import InputWithLabel from './InputWithLabel'
import PlusButton from './PlusButton'

export default function KeyGeneratorForm() {
  const [names, setNames] = useState([''])
  const [generatedValues, setGeneratedValues] = useState<{
    sk: string
    pk: string
    nsec: string
    npub: string
  } | null>(null)
  function appendName() {
    setNames([...names, ''])
    setTimeout(
      () => document.getElementsByName(`name-${names.length}`)[0].focus(),
      100,
    )
  }
  async function onSubmit() {
    const response = await fetch('/api/genkeys', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ words: names }),
    })
    if (response.ok) {
      const genValues = await response.json()
      setGeneratedValues(genValues)
    }
  }
  return (
    <>
      <form
        style={{ display: generatedValues ? 'none' : 'block' }}
        onSubmit={(e) => {
          e.preventDefault()
          appendName()
        }}
      >
        {names.map((_, index) => (
          <InputWithLabel
            key={index}
            name={`name-${index}`}
            label={``}
            onChange={(e) => {
              const newNames = [...names]
              newNames[index] = e.target.value.trim()
              setNames(newNames)
            }}
            onClose={() => {
              const newNames = [...names]
              newNames.splice(index, 1)
              setNames(newNames)
            }}
          />
        ))}
        <PlusButton onClick={appendName} />

        <button onClick={onSubmit} type="button">
          Gerar
        </button>
      </form>
      <div style={{ display: generatedValues ? 'block' : 'none' }}>
        <p>sk: {generatedValues?.sk}</p>
        <p>pk: {generatedValues?.pk}</p>
        <p>nsec: {generatedValues?.nsec}</p>
        <p>npub: {generatedValues?.npub}</p>
        <p>Words: {names.join(' - ')}</p>
      </div>
    </>
  )
}
