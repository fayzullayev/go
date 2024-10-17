import { Input } from "antd";
import { JSXElementConstructor, ReactNode } from "react";
import { Controller, RegisterOptions, useFormContext } from "react-hook-form";
import { CustomInputContainer } from "./style.ts";

interface CustomInputProps {
  name: string;
  title: string;
  rules: RegisterOptions;
  value?: any;
  btnType?: "text" | "password" | "number";
  icon?: ReactNode;
}

function CustomInput({
  name,
  rules,
  title,
  btnType = "text",
  icon,
}: CustomInputProps) {
  const { control } = useFormContext();

  let InputElement: JSXElementConstructor<any>;

  switch (btnType) {
    case "password":
      InputElement = Input.Password;
      break;
    case "number":
      InputElement = Input;
      break;
    case "text":
      InputElement = Input;
      break;
  }

  return (
    <CustomInputContainer>
      <Controller
        control={control}
        name={name}
        rules={rules}
        render={({ field, fieldState: { error } }) => (
          <InputElement
            {...field}
            status={error ? "error" : ""}
            placeholder={title}
            addonAfter={icon}
            size={"large"}
          />
        )}
      />
    </CustomInputContainer>
  );
}

export default CustomInput;
