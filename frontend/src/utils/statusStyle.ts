import { Status } from "proto-gen-web/services/backend/resources";

export const submitStatusClass = (status?: Status): string => {
  switch (status) {
    case Status.WAITING:
    case Status.IN_PROGRESS:
      return "bg-teal-500";
    case Status.SUCCESS:
      return "bg-blue-600";
    case Status.CONNECTION_FAILED:
      return "bg-red-600";
    case Status.TIMEOUT:
      return "bg-yellow-500";
    case Status.VALIDATION_ERROR:
    case Status.INTERNAL_ERROR:
    default:
      return "bg-orange-500";
  }
};

export const taskStatusClass = (status?: Status): string => {
  switch (status) {
    case Status.WAITING:
      return "opacity-70";
    case Status.IN_PROGRESS:
      return "bg-teal-500";
    case Status.SUCCESS:
      return "bg-blue-600";
    case Status.CONNECTION_FAILED:
      return "bg-red-500";
    case Status.TIMEOUT:
      return "bg-yellow-500";
    case Status.VALIDATION_ERROR:
    case Status.INTERNAL_ERROR:
      return "bg-orange-500";
    default:
      return "bg-gray-700 opacity-70";
  }
};

export const statusLabel = (status?: Status): string => {
  switch (status) {
    case Status.WAITING:
      return "Waiting";
    case Status.IN_PROGRESS:
      return "In Progress";
    case Status.SUCCESS:
      return "Success";
    case Status.CONNECTION_FAILED:
      return "Connection Failed";
    case Status.VALIDATION_ERROR:
      return "Validation Error";
    case Status.TIMEOUT:
      return "Timeout";
    case Status.INTERNAL_ERROR:
      return "Internal Error";
    default:
      return "Unknown Error";
  }
};
