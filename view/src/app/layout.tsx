import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Footer from "@/components/organisms/Footer/Footer";
import Header from "@/components/organisms/Header/Header";
import "@/style/global.css";
import { cva } from "@@/styled-system/css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "LGTMeme",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <Header />
        <main className={mainRecipe()}>{children}</main>
        <Footer />
      </body>
    </html>
  );
}

const mainRecipe = cva({
  base: {
    bg: "GHOUST_WHITE",
    maxWidth: "100vw",
    md: { minHeight: "calc(100vh - 200px)" },
  },
});
