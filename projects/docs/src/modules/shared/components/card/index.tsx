import type { Component } from "@builder.io/qwik";
import { Body } from "./card.body";
import { Main } from "./card";
import type { ICardsMainProps } from "./card.d";
import { Heading } from "./card.heading";

interface ICard extends Component<ICardsMainProps> {
  Body: typeof Body;
  Heading: typeof Heading;
}
const Card: ICard = Main as ICard;

Card.Body = Body;
Card.Heading = Heading;

export default Card;
