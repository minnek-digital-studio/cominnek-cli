import { Slot, component$ } from '@builder.io/qwik';
import type { ITextProps } from './text.d';
import Styles from './text.module.scss';

export const Text = component$<ITextProps>((props) => {
  const baseClass = [Styles.text, {[Styles.bold]: props.bold}];
  const _class = [...baseClass, props.class];

    if (props.type === 'span') {
        return (
          <span class={_class}>
            <Slot />
          </span>
        );
    }
    
    return (
      <p class={_class}>
        <Slot />
      </p>
    );
});