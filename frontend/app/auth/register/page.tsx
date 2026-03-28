"use client";

import { useState } from "react";
import { Button } from "@/components/selia/button";
import { Card, CardHeader, CardTitle, CardDescription, CardBody } from "@/components/selia/card";
import { Input } from "@/components/selia/input";
import { Field, FieldLabel } from "@/components/selia/field";
import { Text, TextLink } from "@/components/selia/text";

export default function page() {
  const [pending, setPending] = useState(false);

  return (
    <div className="w-full lg:h-screen flex items-center justify-center p-4">
      <Card className="w-full lg:w-5/12 xl:w-md">
        <CardHeader align="center">
          <CardTitle>Register a new account</CardTitle>
          <CardDescription>Create a new account with your email</CardDescription>
        </CardHeader>
        <CardBody className="flex flex-col gap-5">
          <Field>
            <FieldLabel htmlFor="email">Email</FieldLabel>
            <Input id="email" type="email" placeholder="Enter your email" />
          </Field>
          <Field>
            <FieldLabel htmlFor="password">Password</FieldLabel>
            <Input id="password" type="password" placeholder="Enter your password" />
          </Field>
          <Button
            variant="tertiary"
            block
            size="lg"
            progress={pending}
            onClick={() => {
              setPending(true);
              setTimeout(() => setPending(false), 2000);
            }}>
            Register
          </Button>
        </CardBody>
      </Card>
    </div>
  );
}
