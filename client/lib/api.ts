"use server";

import { getServerToken, setServerTokens, removeServerTokens } from "./auth";

// Base API URL
const API_URL = process.env.API_URL;

// Type for API response
type ApiResponse<T> = {
  data: T;
  status: number;
  ok: boolean;
};

// Generic fetch function with authentication and error handling
export async function fetchApi<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  // Get the token from cookies
  const token = await getServerToken();

  // Prepare headers
  const headers = new Headers(options.headers);

  // Only set Content-Type if it's not FormData and not already set
  if (!(options.body instanceof FormData) && !headers.has("Content-Type")) {
    headers.set("Content-Type", "application/json");
  }

  // Add authorization if token exists
  if (token) {
    headers.set("Authorization", `Bearer ${token}`);
  }

  // Prepare the full URL
  const url = `${API_URL}/${
    endpoint.startsWith("/") ? endpoint.slice(1) : endpoint
  }`;

  try {
    // Make the request
    const response = await fetch(url, {
      ...options,
      headers,
    });

    // Handle 401 Unauthorized - token expired or invalid
    if (response.status === 401) {
      // Try to refresh the token
      const refreshed = await refreshToken();

      // If refresh was successful, retry the request with the new token
      if (refreshed) {
        const newToken = await getServerToken();
        headers.set("Authorization", `Bearer ${newToken}`);

        const retryResponse = await fetch(url, {
          ...options,
          headers,
        });

        const data = await retryResponse.json().catch(() => ({}));
        return {
          data,
          status: retryResponse.status,
          ok: retryResponse.ok,
        };
      } else {
        // If refresh failed, clear tokens and return the 401 response
        await removeServerTokens();
        const data = await response.json().catch(() => ({}));
        return {
          data,
          status: response.status,
          ok: false,
        };
      }
    }

    // Parse the response data
    const data = await response.json().catch(() => ({}));

    return {
      data,
      status: response.status,
      ok: response.ok,
    };
  } catch (error) {
    console.error("API request error:", error);
    return {
      data: {} as T,
      status: 500,
      ok: false,
    };
  }
}

// Helper functions for common HTTP methods
export async function get<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  return fetchApi<T>(endpoint, { ...options, method: "GET" });
}

export async function post<T>(
  endpoint: string,
  data: any,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  const body = data instanceof FormData ? data : JSON.stringify(data);
  return fetchApi<T>(endpoint, {
    ...options,
    method: "POST",
    body,
  });
}

export async function put<T>(
  endpoint: string,
  data: any,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  const body = data instanceof FormData ? data : JSON.stringify(data);
  return fetchApi<T>(endpoint, {
    ...options,
    method: "PUT",
    body,
  });
}

export async function patch<T>(
  endpoint: string,
  data: any,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  const body = data instanceof FormData ? data : JSON.stringify(data);
  return fetchApi<T>(endpoint, {
    ...options,
    method: "PATCH",
    body,
  });
}

export async function del<T>(
  endpoint: string,
  options: RequestInit = {}
): Promise<ApiResponse<T>> {
  return fetchApi<T>(endpoint, { ...options, method: "DELETE" });
}

// Token refresh function
async function refreshToken(): Promise<boolean> {
  try {
    const refreshToken = await getServerToken();

    if (!refreshToken) {
      return false;
    }

    const response = await fetch(`${API_URL}/auth/refresh/`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refresh: refreshToken }),
    });

    if (!response.ok) {
      return false;
    }

    const data = await response.json();

    if (data.access && data.refresh) {
      await setServerTokens(data.access, data.refresh);
      return true;
    }

    return false;
  } catch (error) {
    console.error("Token refresh error:", error);
    return false;
  }
}
