import { z } from "zod";

// Register schema
export const registerSchema = z.object({
  name: z.string().min(2, { message: "Name must be at least 2 characters" }),
  email: z.string().email({ message: "Please enter a valid email address" }),
  password: z
    .string()
    .min(6, { message: "Password must be at least 6 characters" }),
});

export type RegisterFormData = z.infer<typeof registerSchema>;

export const loginSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email address" }),
  password: z
    .string()
    .min(3, { message: "Password must be at least 3 characters" }),
});

export type LoginFormData = z.infer<typeof loginSchema>;
