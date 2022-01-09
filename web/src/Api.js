
const staging = 'https://ikfi3oi9te.execute-api.ca-central-1.amazonaws.com/staging/quotes'
const test = 'http://localhost:3000'

const current = () => process.env.REACT_APP_API_ENV === 'test' ? test : staging

const Api = {
  staging, test, current
}

export default Api
