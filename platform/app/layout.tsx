import type { Metadata } from "next";
import "./globals.css"; import MainWrapper from "@/components/Layout";


export const metadata: Metadata = {
  title: "Canto GPT",
  description: "GEN AI for Canto Protocol",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <MainWrapper>
          {children}
        </MainWrapper>
      </body>
    </html>
  );
}
