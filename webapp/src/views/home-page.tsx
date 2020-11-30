import React from "react";


type Props = {
  updateScenario: (scenario: any) => void,
  saveScenario: () => void
}

const HomePage: React.FC = (props: Props) => {
  const updateScenario = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    props.updateScenario(e.target.value);
  };
  const saveScenario = (e: React.FormEvent) => {
    e.preventDefault();
    props.saveScenario();
  };

  return (
    <div>
      <div>
        <textarea
          onBlur={updateScenario}
        />
      </div>
      <button onClick={saveScenario}>Submit</button>
    </div>
  );
};

export default HomePage;
