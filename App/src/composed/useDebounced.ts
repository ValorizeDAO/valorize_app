import { ref } from "vue"

export default function useDebounced(timeoutCount = 800, fn = (args: any) => args) {
  let timeoutRef : (null | number) = null
  const displayValue = ref("")
  const debouncedValue = ref("")

  const debounceListener = (e: Event) => {
    if (timeoutRef) {
      clearTimeout(timeoutRef)
    }

    // @ts-ignore
    const { value }: string = e.target
    displayValue.value = value
    timeoutRef = setTimeout(() => {
      debouncedValue.value = value
      fn(debouncedValue)
    }, timeoutCount)
  }

  return { debouncedValue, displayValue, debounceListener }
}
