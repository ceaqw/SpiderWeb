<template>
    <div id="trend-fail">
        <Filter :filter="filterForm"/> 
        <h2>Fail</h2>
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
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.failFilter,
        }
    },
}
</script>
