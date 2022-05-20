<template>
    <v-row justify="center">
        <v-dialog v-model="dialog" scrollable>
            <template v-slot:activator="{ props }">
                <v-btn icon="mdi-eyedropper-variant" :color="this.color || '#FFFFFFFF'" v-bind="props" :rounded="0" flat />
            </template>
            <v-card flat>
                <v-card-title>Select color</v-card-title>
                <v-divider></v-divider>
                    <v-color-picker v-model="int_color" mode="hexa" :modes="['hexa']" hide-inputs show-swatches swatches-max-height="150px" />
                <v-divider></v-divider>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue-darken-1" text @click="cancel">Cancel</v-btn>
                    <v-btn color="blue-darken-1" text @click="save">Ok</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-row>
</template>

<script>
export default {
    name: 'color-selector',
    data() {
        return {
            dialog: false,
            int_color: this.color || '#FFFFFFFF'
        }
    },
    props: {
        color: String,
        onColor: Function
    },
    methods: {
        save() {
            this.dialog = false;
            this.onColor(this.int_color);
        },
        cancel() {
            this.dialog = false;
            this.int_color = this.color;
        }
    }
}
</script>

<style>
.v-color-picker.v-sheet {
    border-radius: unset;
    box-shadow: unset;
}

.v-color-picker__controls .v-color-picker-preview {
    margin-bottom: 0px;
}
</style>