/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        back: "url(../public/resourses/mainBack.jpg)",
        banner: "url(../public/resourses/navBar.jpg)",
      }
    },
  },
  plugins: [],
}