<template>
    <div id="project">
        <Filter :filter="filterForm"/> 
        <h2>Project</h2>
    </div>
</template>

<script>
import Filter from '@/components/filter.vue'

export default {
    components: {
        Filter,
    },
    data() {
        return {
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.projectFilter,
        }
    },
}
</script>
