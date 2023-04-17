import { useSignal, useVisibleTask$, $ } from "@builder.io/qwik";

const delay = (time: number) => new Promise((res) => setTimeout(res, time));

export const useTextAnimation = (words: string[]) => {
  const textChars = useSignal("");
  const currentWord = useSignal(0);
  const timeOut = 150;
  const delayTime = timeOut * 15;

  const type = $((word: string) => {
    word.split("").forEach((char, i) => {
      setTimeout(() => {
        textChars.value += char;
      }, timeOut * i);
    });
  });

  const clear = $((word: string) => {
    const len = word.length;
    for (let i = len - 1; i >= 0; i--) {
      setTimeout(() => {
        textChars.value = textChars.value.slice(0, -1);

        if (i === 0 && currentWord.value < words.length - 1) {
          currentWord.value = currentWord.value + 1;
        }
      }, timeOut * (len - i) + delayTime);
    }
  });

  const animate = $(async (word: string) => {
    await type(word);
    await delay(delayTime);
    await clear(word);
    await delay(timeOut * 5);
  });

  useVisibleTask$(async ({ track }) => {
    const index = track(currentWord);
    const word = words[index.value];
    animate(word).then(() => {
      if (index.value === words.length - 1) {
        setTimeout(() => {
          currentWord.value = 0;
        }, delayTime + 500);
      }
    });
  });

  return textChars;
};
