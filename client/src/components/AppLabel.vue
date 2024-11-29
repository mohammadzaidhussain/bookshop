<template>
    <div>
        <slot :styleClass="styleClass" :cssStyles="cssStyles"></slot>
    </div>
</template>

<script>

export default {
    name: 'AppLabel',
    props: {
        styleClass: {
            type:String,
            default: ''
        },
        cssStyles: {
            type: String,
            default: ''
        }
    }
}

</script>