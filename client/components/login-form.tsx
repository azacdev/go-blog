"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

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

const formSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email address" }),
  password: z
    .string()
    .min(6, { message: "Password must be at least 6 characters" }),
});

type FormValues = z.infer<typeof formSchema>;

export default function LoginForm() {
  const [errors, setErrors] = useState<{ Email?: string; Password?: string }>(
    {}
  );
  const router = useRouter();

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: FormValues) {
    console.log("Login form values:", values);

    // Here you would typically handle authentication
    // For now, we'll just log the values

    // Mock API call
    try {
      // Simulate API call
      // await fetch('/api/login', {
      //   method: 'POST',
      //   body: JSON.stringify(values),
      // })
      // If successful, redirect
      // router.push('/dashboard')
    } catch (error) {
      console.error("Login error:", error);
      setErrors({
        Email: "Invalid email or password",
      });
    }
  }

  return (
    <div className="w-full max-w-md mx-auto">
      <div className="bg-white rounded-lg shadow-md overflow-hidden relative">
        <div className="absolute top-0 left-0 w-0 h-0 border-t-[50px] border-t-primary border-r-[50px] border-r-transparent"></div>

        <h2 className="text-2xl font-bold p-6 text-center">Log in</h2>

        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="p-6 pt-0 space-y-4"
          >
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder="Email" {...field} />
                  </FormControl>
                  {errors.Email && (
                    <p className="text-sm text-red-500">{errors.Email}</p>
                  )}
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
                  {errors.Password && (
                    <p className="text-sm text-red-500">{errors.Password}</p>
                  )}
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button type="submit" className="w-full">
              Log in
            </Button>
          </form>
        </Form>

        <p className="text-center pb-6">
          Or{" "}
          <Link href="/register" className="text-primary hover:underline">
            Register
          </Link>
        </p>
      </div>
    </div>
  );
}
