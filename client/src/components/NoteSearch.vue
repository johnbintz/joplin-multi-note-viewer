<script setup>
import { ref, watch, onMounted } from 'vue'

const searchRef = ref(null)
const noteSearchRef = ref(null)

const emit = defineEmits(['search', 'select'])
const props = defineProps({
  results: { type: Array, required: true },
  searchFocus: { type: Number, required: true }
})

const search = ref('')
const activeKBElement = ref(0)

watch(search, (nv) => {
  emit('search', nv)
  noteSearchRef.value.scrollIntoView({ block: 'end', inline: 'end' })
})
watch(() => props.searchFocus, () => {
  noteSearchRef.value.scrollIntoView({ block: 'end', inline: 'end' })
  searchRef.value.focus()
})

function select(id) {
  emit('select', id)
  search.value = ''
  activeKBElement.value = 0
}

onMounted(() => {
  searchRef.value.focus()

  searchRef.value.addEventListener('keydown', (e) => {
    if ([40,38,13].includes(e.keyCode) && props.results.length > 0) {
      e.preventDefault()
      switch (e.keyCode) {
        case 40: //down
          if (activeKBElement.value < props.results.length) {
            activeKBElement.value++
          }
          break;
        case 38: //up
          if (activeKBElement.value > 0) {
            activeKBElement.value--
          }
          break;
        case 13:
          select(props.results[activeKBElement.value].id)
          break;
      }
    }
  })
})
</script>

<template>
  <div class="note-search" ref="noteSearchRef">
    <div class="form-element">
      <label>Search</label>
      <input placeholder="Enter 2 or more characters" ref="searchRef" v-model="search" />
    </div>
    <div
      class="result"
      @click="select(result.id)"
      :class="{'is-active': index === activeKBElement}"
      v-for="(result, index) in results"
    >
      {{ result.title }}
    </div>
  </div>
</template>

<style lang="scss" scoped>
.form-element {
  display: flex;
  align-items: center;

  label {
    font-weight: bold;
    flex: 0 1 auto;
    margin-right: 1rem;
    font-size: 1.5rem;
  }

  input {
    flex: 1 0 auto;
    font-size: 1.5rem;
  }
}

.result {
  border: solid #222 1px;
  padding: 0.5rem;
  cursor: pointer;

  &:hover, &.is-active {
    color: #00f;
    background-color: #ddd;
  }
}
</style>
