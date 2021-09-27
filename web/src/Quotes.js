import { useState, useEffect } from 'react'
import api from './Api.js'

const Quotes = () => {
  const [error, setError] = useState(null)
  const [isLoaded, setIsLoaded] = useState(false)
  const [items, setItems] = useState([])

  useEffect(() => {
    fetch(api.staging)
      .then(res => {
        return res.json()
      })
      .then( res => JSON.parse(res.body) )
      .then(items => {
        setIsLoaded(true)
        setItems(items)
      }, err => {
        setIsLoaded(false)
        setError(err)
      })
  }, [])

  if (error) {
    return <div>Error: {error.message}</div>
  } else if (!isLoaded) {
    return <div>Loading...</div>
  } else {
    return (
      <ul>
        {items.map(item => (
          <li key={item.id}>
            &quot;{item.text}&quot; - {item.author}
          </li>
        ))}
      </ul>
    )
  }
}

export default Quotes