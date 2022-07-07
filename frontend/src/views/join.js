const html = require('choo/html')

const Signup = require('../components/forms/signup')

module.exports = (state, emit) => {
  const signup = state.cache(Signup, 'signup')

  return html`
    <div class="flex flex-column">
      <h2 class="f3 fw1 mt3 near-black near-black--light light-gray--dark lh-title">Join now</h2>
      ${signup.render()}
    </div>
  `
}
