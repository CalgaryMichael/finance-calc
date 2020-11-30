import React from "react";


type Props = {
  testValue: string,
  test: () => void
}

const HomePage: React.FC = (props: Props) => {
  const runTest = (e: React.FormEvent) => {
    e.preventDefault();
    props.test();
  };

  return <button onClick={runTest}>{ props.testValue }</button>;
};

export default HomePage;

