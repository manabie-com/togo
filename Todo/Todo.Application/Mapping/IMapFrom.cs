using AutoMapper;

namespace Todo.Application.Mapping
{
    public interface IMapFrom<T>
    {
        void Mapping(Profile profile);
    }
}
