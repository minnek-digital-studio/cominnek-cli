type IButtonVariants = "primary" | "secondary" | "dark" | "light";

type IButtonBase = {
  variant?: IButtonVariants;
  class?: string;
  title?: string;
};

type IButtonGroupBase = {
  class?: string;
  wrap?: boolean;
};

type IButtonLinkProps = IButtonBase & {
  href: string;
  _target?: "_blank" | "_self" | "_parent" | "_top";
};

type IButtonBtnProps = IButtonBase & {
  href?: undefined;
  _target?: undefined;
  type?: "button" | "submit" | "reset";
};

export type IButtonProps = IButtonLinkProps | IButtonBtnProps;
export type IButtonGroupProps =
  | (IButtonGroupBase & {
      variant?: "inline";
    })
  | (IButtonGroupBase & {
      variant?: "spaced";
      gap?: number;
    });
