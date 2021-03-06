# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from metrics.v3 import ingestion_service_pb2 as metrics_dot_v3_dot_ingestion__service__pb2


class IngestionStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Submit = channel.unary_unary(
                '/kin.agora.metrics.v3.Ingestion/Submit',
                request_serializer=metrics_dot_v3_dot_ingestion__service__pb2.SubmitRequest.SerializeToString,
                response_deserializer=metrics_dot_v3_dot_ingestion__service__pb2.SubmitResponse.FromString,
                )


class IngestionServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Submit(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_IngestionServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Submit': grpc.unary_unary_rpc_method_handler(
                    servicer.Submit,
                    request_deserializer=metrics_dot_v3_dot_ingestion__service__pb2.SubmitRequest.FromString,
                    response_serializer=metrics_dot_v3_dot_ingestion__service__pb2.SubmitResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kin.agora.metrics.v3.Ingestion', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Ingestion(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Submit(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.metrics.v3.Ingestion/Submit',
            metrics_dot_v3_dot_ingestion__service__pb2.SubmitRequest.SerializeToString,
            metrics_dot_v3_dot_ingestion__service__pb2.SubmitResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
