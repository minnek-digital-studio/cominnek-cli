export interface ICardsMainProps {
  /** The maximum width of the card.
   * If set to 0, the card will not have a maximum width.
   * */
  maxWidth?: string;

  /** Whether the card should have a shadow */
  shadow?: boolean;

  /** The variant of the card */
  variant?: "default" | "outlined";

  /** Class to apply to the card */
  class?: string;
}
