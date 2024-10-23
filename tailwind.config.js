/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./view/**/*.templ",
    "./components/*.templ",
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    styled: true,
    themes: [
      {
        mytheme: {
          primary: "#669BBC",
          secondary: "#FDF0D5",
          accent: "#e879f9",
          neutral: "#201e24",
          "base-100": "#FAFAFA",
          info: "#67c6ff",
          success: "#66BC69",
          warning: "#C16612",
          error: "#780000",
        },
      },
    ],
  },
  plugins: [
    require('daisyui'),
  ],
}

