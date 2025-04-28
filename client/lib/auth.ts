"use server";

import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import type { ResponseCookie } from "next/dist/compiled/@edge-runtime/cookies";

// Token constants
const ACCESS_TOKEN = "auth_token";
const REFRESH_TOKEN = "refresh_token";

// Cookie options
const COOKIE_OPTIONS: Partial<ResponseCookie> = {
  // 7 days in seconds
  maxAge: 7 * 24 * 60 * 60,
  secure: process.env.NODE_ENV === "production",
  httpOnly: true, // Important for security - prevents client-side JS access
  path: "/",
  sameSite: "strict", // This needs to be a literal "strict", "lax", or "none"
};

// Server-side token management
export async function getServerToken(): Promise<string | undefined> {
  const cookieStore = await cookies();
  return cookieStore.get(ACCESS_TOKEN)?.value;
}

export async function getServerRefreshToken(): Promise<string | undefined> {
  const cookieStore = await cookies();
  return cookieStore.get(REFRESH_TOKEN)?.value;
}

export async function setServerTokens(
  accessToken: string,
  refreshToken: string
): Promise<void> {
  const cookieStore = await cookies();
  cookieStore.set(ACCESS_TOKEN, accessToken, COOKIE_OPTIONS);
  cookieStore.set(REFRESH_TOKEN, refreshToken, COOKIE_OPTIONS);
}

export async function removeServerTokens(): Promise<void> {
  const cookieStore = await cookies();
  cookieStore.delete(ACCESS_TOKEN);
  cookieStore.delete(REFRESH_TOKEN);
}

// Server-side authentication check
export async function isServerAuthenticated(): Promise<boolean> {
  return !!(await getServerToken());
}

// Auth-related server actions
export async function signOut(): Promise<void> {
  await removeServerTokens();
  redirect("/login");
}
