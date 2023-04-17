interface ISpacerProps {
  size?: number;
}

export const Spacer = ({ size = 20 }: ISpacerProps) => {
  return (
    <div
      style={{
        height: `${size}px`,
      }}
      aria-hidden="true"
    ></div>
  );
};
