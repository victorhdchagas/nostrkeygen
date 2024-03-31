interface InputWithLabelProtocol {
  name: string
  label: string
  type?: string
  id?: string
  placeholder?: string
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void
  onClose?: () => void
}
export default function InputWithLabel({
  name,
  label,
  type = 'text',
  id,
  placeholder,
  onChange,
  onClose,
}: InputWithLabelProtocol) {
  return (
    <div style={{ position: 'relative' }}>
      <label htmlFor={name}>{label}</label>
      <input
        id={id ?? name}
        name={name}
        type={type}
        onChange={onChange}
        placeholder={placeholder}
        style={{ width: '100%', borderRadius: '5px', padding: '5px' }}
      />
      <span
        style={{
          cursor: 'pointer',
          marginLeft: '5px',
          color: '#afafaf',
          fontSize: '10px',
          position: 'absolute',
          top: 8,
          right: 1,
        }}
        onClick={onClose}
      >
        remove
      </span>
    </div>
  )
}
