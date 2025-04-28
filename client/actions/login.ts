"use server";

import { post } from "@/lib/api";
import { LoginFormData, loginSchema } from "@/lib/schema";
import type { LoginResponse, LoginActionResponse } from "@/types/login";

export async function login(
  formData: LoginFormData
): Promise<LoginActionResponse> {
  // Validate the input data
  const validationResult = loginSchema.safeParse(formData);

  if (!validationResult.success) {
    console.log(
      "Validation errors:",
      validationResult.error.flatten().fieldErrors
    );

    return {
      success: false,
      errors: validationResult.error.flatten().fieldErrors,
      message: "Validation failed",
    };
  }

  try {
    const response = await post<LoginResponse>("/login", validationResult.data);

    if (response.ok) {
      // Login successful
      return {
        success: true,
        message:
          typeof response.data.message === "string"
            ? response.data.message
            : "Login successful",
        user: response.data.user,
      };
    } else if (response.status === 401) {
      // Unauthorised
      return {
        success: false,
        message:
          typeof response.data.message === "string"
            ? response.data.message
            : "Invalid credentials",
      };
    } else if (response.status === 400 || response.status === 409) {
      // Validation errors from the server
      const serverErrors: Record<string, string[]> = {};

      if (response.data.message && typeof response.data.message === "object") {
        Object.entries(response.data.message).forEach(([key, value]) => {
          serverErrors[key.toLowerCase()] = Array.isArray(value)
            ? value
            : [value as string];
        });
      }

      return {
        success: false,
        message: "Validation failed",
        errors: serverErrors,
      };
    } else {
      // Other errors
      return {
        success: false,
        message:
          typeof response.data.message === "string"
            ? response.data.message
            : "An error occurred during login",
      };
    }
  } catch (error) {
    console.error("Login error:", error);
    return {
      success: false,
      message: "Failed to connect to the server. Please try again later.",
    };
  }
}
