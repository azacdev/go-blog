"use server";

import { post } from "@/lib/api";
import { type RegisterFormData, registerSchema } from "@/lib/schema";
import type {
  RegisterResponse,
  RegisterActionResponse,
} from "@/types/register";

export async function register(
  formData: RegisterFormData
): Promise<RegisterActionResponse> {
  // Validate the input data
  const validationResult = registerSchema.safeParse(formData);

  if (!validationResult.success) {
    return {
      success: false,
      errors: validationResult.error.flatten().fieldErrors,
      message: "Validation failed",
    };
  }

  try {
    const response = await post<RegisterResponse>(
      "/register",
      validationResult.data
    );

    if (response.ok) {
      // Registration successful
      return {
        success: true,
        message:
          typeof response.data.message === "string"
            ? response.data.message
            : "Registration successful",
      };
    } else if (response.status === 409) {
      // Email already exists
      return {
        success: false,
        message:
          typeof response.data.message === "string"
            ? response.data.message
            : "Email address already exists",
        errors: { email: ["Email address already exists"] },
      };
    } else if (response.status === 400) {
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
            : "An error occurred during registration",
      };
    }
  } catch (error) {
    console.error("Registration error:", error);
    return {
      success: false,
      message: "Failed to connect to the server. Please try again later.",
    };
  }
}
