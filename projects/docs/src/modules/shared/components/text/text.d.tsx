interface ITextPropsBase {
  type?: "p" | "span";
  class?: string;
  bold?: boolean;
}

interface IPProps extends ITextPropsBase {
  type?: "p";
}

interface ISpanProps extends ITextPropsBase {
  type?: "span";
}

export type ITextProps = IPProps | ISpanProps;
