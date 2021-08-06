import { Ref, ref } from "vue";

export default function (timeoutCount = 800, fn = (args: Ref<string>) => args) {
  let timeoutRef = null as null | number;
  const displayValue = ref("");
  const debouncedValue = ref("");

  const debounceListener = (e: InputEvent) => {
    if (timeoutRef !== null) {
      clearTimeout(timeoutRef);
    }

    // @ts-ignore
    const { value }: string = e.target
    displayValue.value = value;
    timeoutRef = setTimeout(() => {
      debouncedValue.value = value;
      fn(debouncedValue);
    }, timeoutCount);
  };

  return { debouncedValue, displayValue, debounceListener };
}
