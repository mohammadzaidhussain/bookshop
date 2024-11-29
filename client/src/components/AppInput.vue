<template>
    <ion-input 
        ref="input" 
        fill="outline" 
        :type="type" 
        :label="label" 
        :label-placement="labelPosition"
        :placeholder="placeholder" 
        :counter="counter" 
        :maxlength="maxLength" 
        :minlength="minLength" 
        :disabled="disabled"
        :class="styleClass" 
        :value="input" 
        :debounce="debounce" 
        :readonly="readonly" 
        :required="required"
        @ionInput="onInput($event)"
        @ionBlur="markTouched($event)"
    />
</template>

<script>

import { IonInput } from '@ionic/vue';

export default {
    name: 'AppInput',
    components: {
        IonInput
    },
    props: {
        type: {
            type: String,
            default: 'text'
        },
        label: {
            type: String,
            default: ''
        },
        labelPosition: {
            type: String,
            default: ''
        },
        placeholder: {
            type: String,
            default: ''
        },
        counter: {
            type: Boolean,
            default: false
        },
        maxLength: {
            type: Number,
            default: 0
        },
        minLength: {
            type: Number,
            default: 0
        },
        debounce: {
            type: Number,
            default: 0
        },
        disabled: {
            type: Boolean,
            default: false
        },
        readonly: {
            type: Boolean,
            default: false
        },
        required: {
            type: Boolean,
            default: false
        },
        input: {
            type: [String, Number],
            default: ''
        },
        styleClass: {
            type: String,
            default: ''
        }
    },
    methods: {
        markTouched(event) {
            if(this.$refs.input.$el) {
                //this.$refs.input.$el?.classList?.add('ion-touched');
            }
        },
        validateEmail(email) {
            if(!email) {
                return false;
            }
            return email.includes("@")
        },
        validatePassword(password) {
            return true;
        },
        validate(val) {
            if(this.$refs.input.$el) {
                // this.$refs.input.$el?.classList?.remove('ion-invalid');   
                // this.$refs.input.$el?.classList?.remove('ion-valid');    
            }
            if(val === '') {
                return false;
            } 
            if(val.toString().length < this.minLength || val.toString().length > this.maxLength) {
                if(this.$refs.input.$el) {
                    // this.$refs.input.$el?.classList?.add('ion-invalid');
                }
                return false;
            }
            else {
                if(this.$refs.input.$el) {
                    // this.$refs.input.$el?.classList?.add('ion-valid');
                }
                return true;
            }
        },
        onInput(event) {
            const val = event.target.value;
            this.$emit('update:input', val);
            if(this.type == "email") {
                const isEmailValidated = this.validateEmail(val);
                this.$emit('on-input', { value: val, isValidated: isEmailValidated, type: this.type });
            }
            else {
                const data = {
                    value: val,
                    isValidated: true,
                    type: this.type
                }
                if(this.type == "password") {
                    data["isValidated"] = this.validatePassword(val);
                }
                if(this.minLength > 0) {
                    data["isValidated"] = this.validate(val)
                }
                this.$emit('on-input', data)
            }
        }
    }
}

</script>