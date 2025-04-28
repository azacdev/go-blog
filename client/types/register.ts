// Auth-related type definitions
export interface RegisterSuccessResponse {
  message: string;
  status: number;
  user?: {
    ID: number;
    Image: string;
    Name: string;
    Email: string;
  };
  [key: string]: any;
}

export interface RegisterErrorResponse {
  error: string;
  message: string | Record<string, string>;
  status: number;
  [key: string]: any;
}

export type RegisterResponse = RegisterSuccessResponse | RegisterErrorResponse;

// Type guard as a simple function
export function isRegisterErrorResponse(
  response: any
): response is RegisterErrorResponse {
  return response && typeof response === "object" && "error" in response;
}

// Response type for the register action
export interface RegisterActionResponse {
  success: boolean;
  message: string;
  errors?: Record<string, string[]>;
  redirect?: boolean;
}
