import { ErrorContainer } from "./style.ts";
import { isRouteErrorResponse, Link, useRouteError } from "react-router-dom";
import { Button, Result } from "antd";
import { ReactNode } from "react";

function Error() {
  const error = useRouteError() as any;

  let errorMsg: ReactNode = "Something went wrong";

  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      errorMsg = <div>This page doesn't exist!</div>;
    }

    if (error.status === 401) {
      errorMsg = <div>You aren't authorized to see this</div>;
    }

    if (error.status === 503) {
      errorMsg = <div>Looks like our API is down</div>;
    }

    if (error.status === 418) {
      errorMsg = <div>🫖</div>;
    }
  } else {
    error.status = 0;
  }

  return (
    <ErrorContainer>
      <Result
        title={error.status ?? 400}
        subTitle={errorMsg}
        extra={
          <Button type="primary">
            <Link to={"/"}>Back Home</Link>
          </Button>
        }
      />
    </ErrorContainer>
  );
}

export default Error;
