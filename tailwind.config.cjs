/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  darkMode: 'class',
  theme: {
      "extend": {
          "colors": {
              "primary": {
                  "dark": {
                      900: "#001a15",
                      800: "#004d3e",
                      700: "#008067",
                      600: "#00b390",
                      500: "#00e6b9",
                      400: "#1affd2",
                      300: "#4dffdc",
                      200: "#80ffe6",
                      100: "#b3fff0",
                      50:  "#e5fffa",
                  },
                  "light": {
                      900: "#001c46",
                      800: "#002e74",
                      700: "#0041a2",
                      600: "#0053d1",
                      500: "#0066ff",
                      400: "#2e81ff",
                      300: "#5d9dff",
                      200: "#8bb9ff",
                      100: "#b9d5ff",
                      50:  "#e8f1ff",
                  }
              },
              "accent": {
                  900: "#030317",
                  800: "#090944",
                  700: "#100f71",
                  600: "#3837d0",
                  500: "#6261ea",
                  400: "#807fea",
                  300: "#9b9aec",
                  200: "#bcbbf6",
                  100: "#bcbbf6",
                  50:  "#e9e8fc",
              },
              "neutral": {
                  900: "#111111",
                  800: "#2d2d2d",
                  700: "#393939",
                  600: "#686868",
                  500: "#808080",
                  400: "#979797",
                  300: "#aeaeae",
                  200: "#c5c5c5",
                  100: "#e3e3e3",
                  50: "#f3f3f3",
              },
          }
      }
  },
  plugins: [],
}
