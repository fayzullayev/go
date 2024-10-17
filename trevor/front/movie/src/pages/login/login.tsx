import { FormProvider, SubmitHandler, useForm } from "react-hook-form";
import { Button } from "antd";
import { FormWrapper, LoginContainer } from "./style.ts";
import { UserOutlined, KeyOutlined } from "@ant-design/icons";
import CustomInput from "../../components/form/custom-input";
import { forwardRef } from "react";

type Fields = {
  login: string;
  password: string;
};

function Login() {
  const methods = useForm<Fields>({ mode: "onSubmit" });
  const { handleSubmit } = methods;

  const onSubmit: SubmitHandler<Fields> = (data) => {
    console.log(data);
  };

  return (
    <LoginContainer>
      <FormProvider {...methods}>
        <FormWrapper>
          <form onSubmit={handleSubmit(onSubmit)}>
            <CustomInput
              title={"Login"}
              name={"login"}
              rules={{ required: true }}
              icon={<UserOutlined />}
            />

            <br />
            <br />

            <CustomInput
              title={"Password"}
              name={"password"}
              rules={{ required: true }}
              btnType={"password"}
              icon={<KeyOutlined />}
            />

            <br />
            <br />
            <Button type="primary" htmlType={"submit"}>
              Primary Button
            </Button>
          </form>
        </FormWrapper>
      </FormProvider>
    </LoginContainer>
  );
}

export default Login;
