import { createDiscreteApi } from "naive-ui"


const { message } = createDiscreteApi(['message']);
const MessageType = ref('')


export function useMessage() {
    function destroyAll() {
      message?.destroyAll();
    }
  
    function create(content, option?) {
      MessageType.value = 'create'
      message?.destroyAll();
      message?.create(content, option);
    }
  
    function error(content, option?) {
        MessageType.value = 'error'
        message?.destroyAll();
        message?.error(content, option);
      
    }
  
    function info(content, option?) {
      MessageType.value = 'info'
      message?.destroyAll();
      message?.info(content, option);
    }
  
    function loading(content, option?) {
      MessageType.value = 'loading'
      message?.destroyAll();
      message?.loading(content, option);
    }
  
    function success(content, option?) {
      MessageType.value = 'success'
      message?.destroyAll();
      message?.success(content, option);
    }
  
    function warning(content, option?) {
      MessageType.value = 'warning'
      message?.destroyAll();
      message?.warning(content, option);
    }
  
    return {
      destroyAll,
      create,
      error,
      info,
      loading,
      success,
      warning,
    };
  }
  