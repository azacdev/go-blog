// Auth-related type definitions
import { User } from "@/types/user";

export interface LoginSuccessResponse {
  message: string;
  status: number;
  user: User;
  [key: string]: any;
}

export interface LoginErrorResponse {
  error: string;
  message: string | Record<string, string>;
  status: number;
  [key: string]: any;
}

export type LoginResponse = LoginSuccessResponse | LoginErrorResponse;

// Type guard as a simple function
export function LoginErrorResponse(
  response: any
): response is LoginErrorResponse {
  return response && typeof response === "object" && "error" in response;
}

// Response type for the register action
export interface LoginActionResponse {
  success: boolean;
  message: string;
  errors?: Record<string, string[]>;
  user?: User;
}
