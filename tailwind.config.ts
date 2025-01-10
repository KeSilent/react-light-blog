import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        primary: "var(--primary)",
        secondary: "var(--secondary)",
        "hover-background": "var(--hover-background)",
        "hover-foreground": "var(--hover-foreground)",
        "man-background": "var(--man-background)",
        "border-colos": "var(--border-colors)",
      },
    },
  },
  plugins: [],
} satisfies Config;
