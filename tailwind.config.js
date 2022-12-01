/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      screens: {
          'sm': '428px',
          // => @media (min-width: 428px) { ... }
        },
      },
  },
  plugins: [],
}