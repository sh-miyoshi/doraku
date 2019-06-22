// Run as Development Mode
let INSIDE_BACKEND_SERVER_URL = "http://localhost:8080"

// Run as Production Mode
if (process.env.NODE_ENV === 'production') {
  INSIDE_BACKEND_SERVER_URL = "https://doraku-241004.appspot.com"
}

export const BACKEND_SERVER_URL = INSIDE_BACKEND_SERVER_URL