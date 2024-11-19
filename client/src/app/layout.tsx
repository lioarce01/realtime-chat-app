"use client";

import localFont from "next/font/local";
import "./globals.css";
import { Provider } from "react-redux";
import { persistor, store } from "@/redux/store";
import { PersistGate } from "redux-persist/integration/react";
import { Auth0Provider } from "@auth0/auth0-react";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const redirectUri =
    typeof window !== "undefined" ? window.location.origin : "";
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <Provider store={store}>
          <PersistGate loading={null} persistor={persistor}>
            <Auth0Provider
              domain={process.env.NEXT_PUBLIC_AUTH0_DOMAIN!}
              clientId={process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID!}
              authorizationParams={{
                redirect_uri: redirectUri,
                audience: process.env.NEXT_PUBLIC_AUTH_AUDIENCE,
                scope: "openid profile email",
              }}
              cacheLocation="localstorage"
            >
              {children}
            </Auth0Provider>
          </PersistGate>
        </Provider>
      </body>
    </html>
  );
}
