import { Ref, ref } from "vue";

export default function (timeoutCount = 800, fn = (args: Ref<string>) => args) {
  let timeoutRef = null as null | number;
  const displayValue = ref("");
  const debouncedValue = ref("");

  const debounceListener = (e: { target: { value: string } }) => {
    if (timeoutRef !== null) {
      clearTimeout(timeoutRef);
    }
    displayValue.value = e.target.value;
    timeoutRef = setTimeout(() => {
      debouncedValue.value = e.target.value;
      fn(debouncedValue);
    }, timeoutCount);
  };

  return { debouncedValue, displayValue, debounceListener };
}
