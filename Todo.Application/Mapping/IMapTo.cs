using AutoMapper;

namespace Todo.Application.Mapping
{
    public interface IMapTo<T>
    {
        void Mapping(Profile profile);
    }
}
