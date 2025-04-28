"use client";

import Link from "next/link";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { AlertCircle, CheckCircle2 } from "lucide-react";

import { register } from "@/actions/register";
import { RegisterFormData, registerSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Alert, AlertDescription } from "@/components/ui/alert";

export default function RegisterForm() {
  const router = useRouter();
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [formStatus, setFormStatus] = useState<{
    type: "success" | "error" | null;
    message: string | null;
  }>({
    type: null,
    message: null,
  });

  const form = useForm<RegisterFormData>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: RegisterFormData) {
    setIsSubmitting(true);
    setFormStatus({ type: null, message: null });

    try {
      const result = await register(values);

      if (result.success) {
        setFormStatus({
          type: "success",
          message:
            typeof result.message === "string"
              ? result.message
              : "Registration successful! Redirecting to login...",
        });

        // Redirect to login page after a short delay
        setTimeout(() => {
          router.push("/login");
        }, 1500);
      } else {
        // Handle validation errors from the server
        if (result.errors) {
          Object.entries(result.errors).forEach(([field, messages]) => {
            if (field in form.formState.errors) {
              form.setError(field as keyof RegisterFormData, {
                type: "server",
                message: Array.isArray(messages) ? messages[0] : messages,
              });
            }
          });
        }

        setFormStatus({
          type: "error",
          message:
            typeof result.message === "string"
              ? result.message
              : "Registration failed. Please try again.",
        });
      }
    } catch (error) {
      console.error("Registration error:", error);
      setFormStatus({
        type: "error",
        message: "An unexpected error occurred. Please try again later.",
      });
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="w-full max-w-md mx-auto">
      <div className="bg-white rounded-lg shadow-md overflow-hidden relative">
        <div className="absolute top-0 left-0 w-0 h-0 border-t-[50px] border-t-primary border-r-[50px] border-r-transparent"></div>

        <h2 className="text-2xl font-bold p-6 text-center">Register</h2>

        {formStatus.message && (
          <Alert
            className={`mx-6 ${
              formStatus.type === "success" ? "bg-green-50" : "bg-red-50"
            }`}
          >
            {formStatus.type === "success" ? (
              <CheckCircle2 className="h-4 w-4 text-green-600" />
            ) : (
              <AlertCircle className="h-4 w-4 text-red-600" />
            )}
            <AlertDescription
              className={
                formStatus.type === "success"
                  ? "text-green-700"
                  : "text-red-700"
              }
            >
              {formStatus.message}
            </AlertDescription>
          </Alert>
        )}

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="p-6 pt-0 space-y-4"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder="Name" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder="Email" type="email" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input type="password" placeholder="Password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              type="submit"
              className="w-full hover:cursor-pointer"
              disabled={isSubmitting}
            >
              {isSubmitting ? "Registering..." : "Register"}
            </Button>
          </form>
        </Form>

        <p className="text-center pb-6">
          Or{" "}
          <Link href="/login" className="text-primary hover:underline">
            Login
          </Link>
        </p>
      </div>
    </div>
  );
}
